@aws
Feature: Service apply

  Scenario: Applying a basic service
    Given I setup ernest with target "https://ernest.local"
    And I'm logged in as "usr" / "pwd"
    And I start recording
    And I apply the definition "aws1.yml"
    And I stop recording
    And an event "network.create.aws-fake" should be called exactly "1" times
    And an event "instance.create.aws-fake" should be called exactly "1" times
    And an event "firewall.create.aws-fake" should be called exactly "1" times
    And all "network.create.aws-fake" messages should contain a field "_type" with "aws-fake"
    And all "network.create.aws-fake" messages should contain a field "datacenter_region" with "fake"
    And all "network.create.aws-fake" messages should contain a field "vpc_id" with "fakeaws"
    And all "network.create.aws-fake" messages should contain a field "range" with "10.1.0.0/24"
    And all "firewall.create.aws-fake" messages should contain a field "vpc_id" with "fakeaws"
    And all "firewall.create.aws-fake" messages should contain a field "datacenter_region" with "fake"
