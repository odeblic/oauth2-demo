package com.nalfiro;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.servlet.view.RedirectView;

@Controller
public class Server {

    private final String ourClientID = "app-0";
    private final String ourClientSecret = "secret-000";
    private final Client client;
    private final Requests requests = new Requests();

    @Autowired
    public Server(Client client) {
        this.client = client;
    }

    @GetMapping("/")
    public RedirectView root() {
        return new RedirectView("/index");
    }

    @GetMapping("/index")
    public String index(Model model) {
        String state = this.requests.requestAuthorization();
        model.addAttribute("state", state);
        return "index";
    }

    @GetMapping("/view")
    public Object view(Model model, @RequestParam String code, @RequestParam String state) {
        Boolean granted = this.requests.grantAuthorization(state, code);

        if (!granted) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Mismatching state!");
        }

        String token = client.fetchToken(ourClientID, ourClientSecret, code);
        String secret = client.fetchSecret(token);
        model.addAttribute("content", secret);
        return "view";
    }
}
