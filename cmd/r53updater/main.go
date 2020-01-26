package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func route53Update(sess *session.Session, zone, fqdn, ip string) {
	r53 := route53.New(sess)

	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(fqdn),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(ip),
							},
						},
						TTL:  aws.Int64(300),
						Type: aws.String("A"),
					},
				},
			},
			Comment: aws.String("auto updating at " + time.Now().String()),
		},
		HostedZoneId: aws.String(zone),
	}

	output, err := r53.ChangeResourceRecordSets(input)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(output)
}

func main() {
	cfg := NewConfig()

	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	// check credentials were loaded
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	}

	tick := time.Tick(time.Duration(cfg.tick) * time.Second)

	ip := NewIPIfy()
	ipPrevious := ""

	for {
		log.Println(time.Now().String())

		ipCurrent, err := ip.Fetch()
		if err != nil {
			continue
		}

		if ipCurrent != ipPrevious {
			log.Printf("ip has changed from %s to %s\n", ipPrevious, ipCurrent)

			ipPrevious = ipCurrent
			route53Update(sess, cfg.zone, cfg.fqdn, ipCurrent)
		}

		select {
		case <-tick:
			continue
		}
	}
}
