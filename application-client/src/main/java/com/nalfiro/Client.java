package com.nalfiro;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.web.client.RestTemplate;

import com.fasterxml.jackson.databind.ObjectMapper;

@Service
public class Client {

    private static class TokenMapper {
        public String access_token;
        public String token_type;
        public int expires_in;
    }

    private static class SecretMapper {
        public String message;
    }

    private final RestTemplate restTemplate;
    private final ObjectMapper objectMapper = new ObjectMapper();

    @Autowired
    public Client(RestTemplate restTemplate) {
        this.restTemplate = restTemplate;
    }

    public String fetchToken(String clientId, String clientSecret, String authorizationCode) {
        String url = "https://localhost:5002/token";

        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_FORM_URLENCODED);

        MultiValueMap<String, String> body = new LinkedMultiValueMap<>();

        body.add("grant_type", "authorization_code");
        body.add("client_id", clientId);
        body.add("client_secret", clientSecret);
        body.add("authorization_code", authorizationCode);

        HttpEntity<MultiValueMap<String, String>> entity = new HttpEntity<>(body, headers);

        ResponseEntity<String> response = restTemplate.exchange(url, HttpMethod.POST, entity, String.class);

        try {
            String content = response.getBody();
            TokenMapper data = objectMapper.readValue(content, TokenMapper.class);
            return data.access_token;
        } catch (Exception e) {
            e.printStackTrace();
            return "";
        }
    }

    public String fetchSecret(String accessToken) {
        String url = "https://localhost:5003/resource";

        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer " + accessToken);

        HttpEntity<String> entity = new HttpEntity<String>(headers);

        ResponseEntity<String> response = restTemplate.exchange(url, HttpMethod.GET, entity, String.class);

        try {
            String content = response.getBody();
            SecretMapper data = objectMapper.readValue(content, SecretMapper.class);
            return data.message;
        } catch (Exception e) {
            e.printStackTrace();
            return "";
        }
    }
}
