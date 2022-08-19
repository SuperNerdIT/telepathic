Feature: server
    In order to work properly
    As a client library
    I need to be able to know the health of server

    Scenario: Call the health endpoint
        Given the main server
        When I call /health endpoint
        Then this returns 200 status code