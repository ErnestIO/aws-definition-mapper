@datacenter @datacenter_update
Feature: Update service with changed datacenter credentials

  Scenario: Non logged user listing
    Given I setup ernest with target "https://ernest.local"
    And I'm logged in as "usr" / "pwd"
    And the datacenter "test_dc" does not exist
    And I run ernest with "datacenter create aws --secret_access_key tmp_secret_access_key --access_key_id tmp_secret_up_to_16_chars --region tmp_region --fake update_datacenter"
    And I run ernest with "service apply definitions/update_datacenter_1.yml"
    And I run ernest with "datacenter update aws --secret_access_key tmp_secret_access_key_2 --access_key_id tmp_secret_up_to_16_chars_2  update_datacenter"
    When I run ernest with "service apply definitions/update_datacenter_2.yml"
    Then "instances" should have been created with field "aws_access_key_id" as "tmp_secret_up_to_16_chars_2"
    And "instances" should have been created with field "aws_secret_access_key" as "tmp_secret_access_key_2"

