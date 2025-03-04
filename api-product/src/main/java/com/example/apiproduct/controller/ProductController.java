package com.example.apiproduct.controller;

import com.example.apiproduct.dto.ProductWriteDto;
import com.example.apiproduct.model.Product;
import com.example.apiproduct.repository.ProductRepository;
import lombok.extern.java.Log;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api")
@Slf4j
public class ProductController {

    private final ProductRepository productRepository;

    public ProductController(ProductRepository productRepository) {
        this.productRepository = productRepository;
    }

    @GetMapping("/products")
    public List<Product> getProducts() {
        log.info("get all products");
        return productRepository.findAll();
    }

    @PostMapping("/product")
    public ResponseEntity<Product> addProduct(@RequestBody ProductWriteDto productDto) {
        Product product = Product.builder()
                .name(productDto.getName())
                .price(productDto.getPrice())
                .build();

        Product savedProduct = productRepository.save(product);
        return new ResponseEntity<>(savedProduct, HttpStatus.CREATED);
    }
}
