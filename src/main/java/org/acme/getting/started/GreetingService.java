package org.acme.getting.started;

import jakarta.enterprise.context.ApplicationScoped;

@ApplicationScoped
public class GreetingService {

    public String greeting(String name) {
        return "hello " + name;
    }
    
    // New method for French greeting
    public String frenchGreeting(String name) {
        return "bonjour " + name;
    }
}
