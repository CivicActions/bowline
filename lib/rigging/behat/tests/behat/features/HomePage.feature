@api @smoke
Feature: Home page
  As a user viewing the home page,
  I want to see a header, footer and other page features
  So that I know I am on the correct site and know how to navigate.

  Scenario: Header page elements
    Given I am on the homepage
     Then I should see the text "Powered by Drupal"
