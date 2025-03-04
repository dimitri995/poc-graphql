package com.example.apiproduct.dto;

import com.example.apiproduct.model.Product;
import lombok.Getter;
import lombok.Setter;
import lombok.Value;

import java.io.Serializable;

/**
 * DTO for {@link Product}
 */
@Value
public class ProductWriteDto implements Serializable {
    String name;
    Double price;
}