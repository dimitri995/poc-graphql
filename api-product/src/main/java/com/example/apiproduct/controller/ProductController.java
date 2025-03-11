package com.example.apiproduct.controller;

import com.example.apiproduct.dto.ProductWriteDto;
import com.example.apiproduct.model.Product;
import com.example.apiproduct.repository.ProductRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@Slf4j
@RequestMapping("/api/products")
public class ProductController {

    private final ProductRepository productRepository;

    public ProductController(ProductRepository productRepository) {
        this.productRepository = productRepository;
    }

    @QueryMapping("allProducts")
    public List<Product> allProducts() {
        return productRepository.findAll();
    }

    @MutationMapping
    public Product addProduct(@Argument ProductWriteDto product) {
        Product productToSave = Product.builder()
                .name(product.getName())
                .price(product.getPrice())
                .build();

        return productRepository.save(productToSave);
    }
}
