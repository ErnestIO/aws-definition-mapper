/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func int64p(i int64) *int64 {
	return &i
}

func TestRDSClusterValidate(t *testing.T) {

	Convey("Given an rds cluster", t, func() {
		nws := []Network{Network{Name: "test-nw"}}
		sgs := []SecurityGroup{SecurityGroup{Name: "test-sg"}}

		r := RDSCluster{
			Name:          "test",
			Engine:        "aurora",
			EngineVersion: "17",
			Port:          int64p(3306),
			AvailabilityZones: []string{
				"eu-west-1",
			},
			SecurityGroups: []string{
				"test-sg",
			},
			Networks: []string{
				"test-nw",
			},
			DatabaseName:     "test",
			DatabaseUsername: "test",
			DatabasePassword: "testtest",
			Backups: RDSBackup{
				Window:    "Mon:22:00-Mon:23:00",
				Retention: int64p(1),
			},
			MaintenanceWindow: "Mon:22:00-Mon:23:00",
			ReplicationSource: "arn:aws:rds:us-east-1:123456789012:cluster:my-aurora-cluster",
			FinalSnapshot:     true,
		}

		Convey("With an invalid name", func() {
			r.Name = ""
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster name should not be null")
				})
			})
		})

		Convey("With a name that exceeds the maximum length", func() {
			r.Name = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster name should not exceed 255 characters")
				})
			})
		})

		Convey("With an invalid engine type", func() {
			r.Engine = ""
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster engine type should not be null")
				})
			})
		})

		Convey("With an invalid database name", func() {
			r.DatabaseName = ""
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database name should not be null")
				})
			})
		})

		Convey("With a database name that exeeds the maximum length", func() {
			r.DatabaseName = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database name should not exceed 64 characters")
				})
			})
		})

		Convey("With a database name that contains disallowed characters", func() {
			r.DatabaseName = "test-1"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database name can only contain alphanumeric characters")
				})
			})
		})

		Convey("With an invalid database username", func() {
			r.DatabaseUsername = ""
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database username should not be null")
				})
			})
		})

		Convey("With a database username that exeeds the maximum length", func() {
			r.DatabaseUsername = "xxxxxxxxxxxxxxxxx"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database username should not exceed 16 characters")
				})
			})
		})

		Convey("With an invalid database password", func() {
			r.DatabasePassword = ""
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database password should not be null")
				})
			})
		})

		Convey("With a database password that does not meed the minimum length", func() {
			r.DatabasePassword = "xxxx"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database password should be between 8 and 41 characters")
				})
			})
		})

		Convey("With a database password that exeeds the maximum length", func() {
			r.DatabasePassword = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database password should be between 8 and 41 characters")
				})
			})
		})

		Convey("With a database password that contains disallowed characters", func() {
			r.DatabasePassword = `¯\_(ツ)_/¯`
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster database password contains an offending character: '¯'")
				})
			})
		})

		Convey("With a port number lower than the minimum allowed", func() {
			r.Port = int64p(1)
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster port number should be between 1150 and 65535")
				})
			})
		})

		Convey("With a port number higher than the maximum allowed", func() {
			r.Port = int64p(999999)
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster port number should be between 1150 and 65535")
				})
			})
		})

		Convey("With a backup retention period lower than the minimum allowed", func() {
			*r.Backups.Retention = 0
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster backup retention should be between 1 and 35 days")
				})
			})
		})

		Convey("With a backup retention period higher than the maximum allowed", func() {
			*r.Backups.Retention = 99
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster backup retention should be between 1 and 35 days")
				})
			})
		})

		Convey("With an invalid backup window", func() {
			r.Backups.Window = "test"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster backup window: Window format must take the form of 'ddd:hh24:mi-ddd:hh24:mi'. i.e. 'Mon:21:30-Mon:22:00'")
				})
			})
		})

		Convey("With an invalid backup window - day", func() {
			r.Backups.Window = "XXX:22:00-Mon:23:00"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster backup window: Date format invalid. Day must be one of Mon, Tue, Wed, Thu, Fri, Sat, Sun")
				})
			})
		})

		Convey("With an invalid backup window - hour", func() {
			r.Backups.Window = "Mon:24:00-Mon:23:00"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster backup window: Date format invalid. Hour must be between 0 and 23 hours")
				})
			})
		})

		Convey("With an invalid backup window - minute", func() {
			r.Backups.Window = "Mon:22:70-Mon:23:00"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster backup window: Date format invalid. Minute must be between 0 and 59 minutes")
				})
			})
		})

		Convey("With an invalid maintenance window", func() {
			r.MaintenanceWindow = "test"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster maintenance window: Window format must take the form of 'ddd:hh24:mi-ddd:hh24:mi'. i.e. 'Mon:21:30-Mon:22:00'")
				})
			})
		})

		Convey("With an invalid maintenance window - day", func() {
			r.MaintenanceWindow = "XXX:22:00-Mon:23:00"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster maintenance window: Date format invalid. Day must be one of Mon, Tue, Wed, Thu, Fri, Sat, Sun")
				})
			})
		})

		Convey("With an invalid maintenance window - hour", func() {
			r.MaintenanceWindow = "Mon:24:00-Mon:23:00"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster maintenance window: Date format invalid. Hour must be between 0 and 23 hours")
				})
			})
		})

		Convey("With a security group that does not exist", func() {
			r.SecurityGroups = []string{"fake"}
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster security group 'fake' does not exist")
				})
			})
		})

		Convey("With a network that does not exist", func() {
			r.Networks = []string{"fake"}
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster network 'fake' does not exist")
				})
			})
		})

		Convey("With an invalid replication source", func() {
			r.ReplicationSource = "fake"
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Cluster replication source should be a valid amazon resource name (ARN), i.e. 'arn:aws:rds:us-east-1:123456789012:cluster:my-aurora-cluster'")
				})
			})
		})

		Convey("With valid fields", func() {
			Convey("When validating the rds cluster", func() {
				err := r.Validate(nws, sgs)
				Convey("Then should not return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

	})
}
