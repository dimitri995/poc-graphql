package com.example.apiuser.dto;

import lombok.Value;

import java.io.Serializable;

/**
 * DTO for {@link com.example.apiuser.model.User}
 */
@Value
public class UserDto implements Serializable {
    String firstname;
    String lastname;
    String businessUnit;
}