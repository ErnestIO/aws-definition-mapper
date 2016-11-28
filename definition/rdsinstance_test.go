/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRDSInstanceValidate(t *testing.T) {

	Convey("Given an rds instance", t, func() {
		nws := []Network{Network{Name: "test-nw"}}
		sgs := []SecurityGroup{SecurityGroup{Name: "test-sg"}}
		cls := []RDSCluster{RDSCluster{Name: "test", Engine: "aurora", EngineVersion: "17", Port: int64p(3306), Networks: []string{"test-nw"}}}

		r := RDSInstance{
			Name:             "test",
			Size:             "db.r3.large",
			Engine:           "aurora",
			EngineVersion:    "17",
			Public:           false,
			Port:             int64p(3306),
			AvailabilityZone: "eu-west-1a",
			HotStandby:       false,
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
			ReplicationSource: "test",
			FinalSnapshot:     true,
		}

		Convey("With an invalid name", func() {
			r.Name = ""
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance name should not be null")
				})
			})
		})

		Convey("With a name that exceeds the maximum length", func() {
			r.Name = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance name should not exceed 255 characters")
				})
			})
		})

		Convey("With an invalid engine type", func() {
			r.Engine = ""
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance engine type should not be null")
				})
			})
		})

		Convey("With an invalid database name", func() {
			r.DatabaseName = ""
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database name should not be null")
				})
			})
		})

		Convey("With a database name that exeeds the maximum length", func() {
			r.DatabaseName = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database name should not exceed 64 characters")
				})
			})
		})

		Convey("With a database name that contains disallowed characters", func() {
			r.DatabaseName = "test-1"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database name can only contain alphanumeric characters")
				})
			})
		})

		Convey("With an invalid database username", func() {
			r.DatabaseUsername = ""
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database username should not be null")
				})
			})
		})

		Convey("With a database username that exeeds the maximum length", func() {
			r.DatabaseUsername = "xxxxxxxxxxxxxxxxx"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database username should not exceed 16 characters")
				})
			})
		})

		Convey("With an invalid database password", func() {
			r.DatabasePassword = ""
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database password should not be null")
				})
			})
		})

		Convey("With a database password that does not meed the minimum length", func() {
			r.DatabasePassword = "xxxx"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database password should be between 8 and 41 characters")
				})
			})
		})

		Convey("With a database password that exeeds the maximum length", func() {
			r.DatabasePassword = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database password should be between 8 and 41 characters")
				})
			})
		})

		Convey("With a database password that contains disallowed characters", func() {
			r.DatabasePassword = `¯\_(ツ)_/¯`
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database password contains an offending character: '¯'")
				})
			})
		})

		Convey("With a port number lower than the minimum allowed", func() {
			r.Port = int64p(1)
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance port number should be between 1150 and 65535")
				})
			})
		})

		Convey("With a port number higher than the maximum allowed", func() {
			r.Port = int64p(999999)
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance port number should be between 1150 and 65535")
				})
			})
		})

		Convey("With a backup retention period lower than the minimum allowed", func() {
			r.Backups.Retention = int64p(0)
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance backup retention should be between 1 and 35 days")
				})
			})
		})

		Convey("With a backup retention period higher than the maximum allowed", func() {
			r.Backups.Retention = int64p(99)
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance backup retention should be between 1 and 35 days")
				})
			})
		})

		Convey("With an invalid backup window", func() {
			r.Backups.Window = "test"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance backup window: Window format must take the form of 'ddd:hh24:mi-ddd:hh24:mi'. i.e. 'Mon:21:30-Mon:22:00'")
				})
			})
		})

		Convey("With an invalid backup window - day", func() {
			r.Backups.Window = "XXX:22:00-Mon:23:00"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance backup window: Date format invalid. Day must be one of Mon, Tue, Wed, Thu, Fri, Sat, Sun")
				})
			})
		})

		Convey("With an invalid backup window - hour", func() {
			r.Backups.Window = "Mon:24:00-Mon:23:00"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance backup window: Date format invalid. Hour must be between 0 and 23 hours")
				})
			})
		})

		Convey("With an invalid backup window - minute", func() {
			r.Backups.Window = "Mon:22:70-Mon:23:00"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance backup window: Date format invalid. Minute must be between 0 and 59 minutes")
				})
			})
		})

		Convey("With an invalid maintenance window", func() {
			r.MaintenanceWindow = "test"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance maintenance window: Window format must take the form of 'ddd:hh24:mi-ddd:hh24:mi'. i.e. 'Mon:21:30-Mon:22:00'")
				})
			})
		})

		Convey("With an invalid maintenance window - day", func() {
			r.MaintenanceWindow = "XXX:22:00-Mon:23:00"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance maintenance window: Date format invalid. Day must be one of Mon, Tue, Wed, Thu, Fri, Sat, Sun")
				})
			})
		})

		Convey("With an invalid maintenance window - hour", func() {
			r.MaintenanceWindow = "Mon:24:00-Mon:23:00"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance maintenance window: Date format invalid. Hour must be between 0 and 23 hours")
				})
			})
		})

		Convey("With a security group that does not exist", func() {
			r.SecurityGroups = []string{"fake"}
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance security group 'fake' does not exist")
				})
			})
		})

		Convey("With a network that does not exist", func() {
			r.Networks = []string{"fake"}
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance network 'fake' does not exist")
				})
			})
		})

		Convey("With no networks and public set to false", func() {
			r.Networks = []string{}
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance should specify at least one network if not set to public")
				})
			})
		})

		Convey("With an invalid license type", func() {
			r.Cluster = ""
			r.Engine = "mysql"
			r.License = "fake"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance license must be one of 'license-included', 'bring-your-own-license', 'general-public-license'")
				})
			})
		})

		Convey("With an invalid storage type", func() {
			r.Cluster = ""
			r.Engine = "mysql"
			r.Storage.Type = "fake"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance storage type must be either 'standard', 'gp2' or 'io1'")
				})
			})
		})

		Convey("With an invalid storage size of less than 5", func() {
			r.Cluster = ""
			r.Engine = "mysql"
			r.Storage.Type = "standard"
			r.Storage.Size = int64p(1)
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance storage size must be between 5 - 6144 GB")
				})
			})
		})

		Convey("With an invalid storage size of greater than 6144", func() {
			r.Cluster = ""
			r.Engine = "mysql"
			r.Storage.Type = "standard"
			r.Storage.Size = int64p(7000)
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance storage size must be between 5 - 6144 GB")
				})
			})
		})

		Convey("With storage iops that are not a multiple of 1000", func() {
			r.Cluster = ""
			r.Engine = "mysql"
			r.Storage.Type = "standard"
			r.Storage.Size = int64p(100)
			r.Storage.Iops = int64p(10500)
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance storage iops must be a multiple of 1000")
				})
			})
		})

		Convey("With a cluster identifier that does not exist", func() {
			r.Cluster = "fake"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance cluster identifier 'fake' does not exist")
				})
			})
		})

		Convey("With an invalid size", func() {
			r.Size = ""
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance size should not be null")
				})
			})
		})

		Convey("With both a cluster and database name set", func() {
			r.Cluster = "test"
			r.DatabaseName = "test"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database name should be set on cluster")
				})
			})
		})

		Convey("With both a cluster and database username set", func() {
			r.Cluster = "test"
			r.DatabaseName = ""
			r.DatabasePassword = ""
			r.DatabaseUsername = "test"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database username should be set on cluster")
				})
			})
		})

		Convey("With both a cluster and database password set", func() {
			r.Cluster = "test"
			r.DatabaseName = ""
			r.DatabaseUsername = ""
			r.DatabasePassword = "test"
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance database password should be set on cluster")
				})
			})
		})

		Convey("With both a cluster and port set", func() {
			r.Cluster = "test"
			r.Engine = ""
			r.EngineVersion = ""
			r.DatabaseName = ""
			r.DatabaseUsername = ""
			r.DatabasePassword = ""
			r.Port = int64p(9999)
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "RDS Instance port should be set on cluster")
				})
			})
		})

		Convey("With attributes defined only on the cluster", func() {
			r = RDSInstance{
				Name:    "test",
				Size:    "db.r3.large",
				Cluster: "test",
			}
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should not return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("With valid fields", func() {
			Convey("When validating the rds instance", func() {
				err := r.Validate(nws, sgs, cls)
				Convey("Then should not return an error", func() {
					So(err, ShouldBeNil)
				})
			})
		})

	})
}
