# Go Boot

*a framework to start a web application quickly like as spring-boot*

![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)

## INTRO

Spring Boot helps you to create Spring-powered, production-grade applications and services with absolute minimum fuss. It takes an opinionated view of the Spring platform so that new and existing users can quickly get to the bits they need.

You can use Spring Boot to create stand-alone Java applications that can be started using java -jar or more traditional WAR deployments. We also provide a command-line tool that runs Spring scripts..

Our primary goals are:

- Provide a radically faster and widely accessible getting started experience for all Spring development.

- Be opinionated, but get out of the way quickly as requirements start to diverge from the defaults.

- Provide a range of non-functional features common to large classes of projects (for example, embedded servers, security, metrics, health checks, externalized configuration).

- Absolutely no code generation and no requirement for XML configuration.

## Installation and Getting Started

The reference documentation includes detailed installation instructions as well as a comprehensive getting started guide.

Here is a quick teaser of a complete Spring Boot application in Java:

Code example:
```
import org.springframework.boot.*;
import org.springframework.boot.autoconfigure.*;
import org.springframework.web.bind.annotation.*;

@RestController
@SpringBootApplication
public class Example {

	@RequestMapping("/")
	String home() {
		return "Hello World!";
	}

	public static void main(String[] args) {
		SpringApplication.run(Example.class, args);
	}

}
```

## Getting Help


## Modules


There are several modules in Spring Boot. Here is a quick overview:

### go-boot

### go-boot-autoconfigure

### go-boot-starters


## LICENCE

Apache License 2.0

