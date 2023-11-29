package com.study.spring5.step02;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.GenericGroovyApplicationContext;
import org.springframework.context.support.GenericXmlApplicationContext;

public class Step02Application {

    public static void main(String[] args) {

        /**
         * AnnotationConfigApplicationContext
         * GenericXmlApplicationContext
         * GenericGroovyApplicationContext
         * */
        AnnotationConfigApplicationContext ctx =
                new AnnotationConfigApplicationContext(AppContext.class);

        Greeter g = ctx.getBean("greeter", Greeter.class);
        String msg = g.greet("스프링");
        System.out.println(msg);
        ctx.close();
    }
}
