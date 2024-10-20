package com.nalfiro;

import java.security.SecureRandom;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;

public class Requests {

    private static String generateState() {
        SecureRandom secureRandom = new SecureRandom();
        byte[] randomBytes = new byte[8];
        secureRandom.nextBytes(randomBytes);

        StringBuilder hexString = new StringBuilder();

        for (byte b : randomBytes) {
            String hex = String.format("%02x", b);
            hexString.append(hex);
        }

        return hexString.toString();
    }

    /*
    public class AuthorizationRequest {
        private long timestamp;
        private String state;
        private String code;

        public AuthorizationRequest(String authorizationCode) {
            this.timestamp = System.nanoTime();
            this.state = "12345";
            this.code = authorizationCode;
        }

        public Boolean isSameState(String state) {
            return state == this.state;
        }

        public Boolean hasExpired() {
            long ellapsedTime = System.nanoTime() - this.timestamp;
            return ellapsedTime >= 1e9 * 60;
        }

        public void setAuthorizationCode(String code) {
            code = this.code;
        }

        public String getAuthorizationCode() {
            return this.code;
        }
    }
    */

    public String requestAuthorization() {
        String state = generateState();
        requests.put(state, null);
        return state;
    }

    public Boolean grantAuthorization(String state, String code) {
        if (requests.keySet().contains(state)) {
            requests.put(state, code);
            return true;
        } else {
            return false;
        }
    }

    public void registerAccessToken(String clientId, String accessToken) {
        this.tokens.get(clientId).add(accessToken);
    }

    //private final Map<String, AuthorizationRequest> requests = new HashMap<>();
    private final Map<String, String> requests = new HashMap<>();
    private final Map<String, Set<String>> tokens = new HashMap<>();
}
