package com.example.apiuser.controller;

import com.example.apiuser.dto.UserWriteDto;
import com.example.apiuser.model.User;
import com.example.apiuser.repository.UserRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api")
@Slf4j
public class UserController {

    private final UserRepository userRepository;

    public UserController(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @GetMapping("/users")
    public List<User> getUsers() {
        return userRepository.findAll();
    }

    @PostMapping("/user")
    public ResponseEntity<User> addUser(@RequestBody UserWriteDto userWriteDto) {
        log.info(String.valueOf(userWriteDto));
        User user = User.builder()
                .firstName(userWriteDto.getFirstname())
                .lastName(userWriteDto.getLastname())
                .businessUnit(userWriteDto.getBusinessUnit())
                .build();

        User savedUser = userRepository.save(user);
        return new ResponseEntity<>(savedUser, HttpStatus.CREATED);
    }
}
