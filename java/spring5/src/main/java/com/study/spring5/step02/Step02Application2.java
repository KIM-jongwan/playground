package com.study.spring5.step02;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class Step02Application2 {

    public static void main(String[] args) {
        AnnotationConfigApplicationContext ctx =
                new AnnotationConfigApplicationContext(AppContext.class);

        Greeter g1 = ctx.getBean("greeter", Greeter.class);
        Greeter g2 = ctx.getBean("greeter", Greeter.class);
        System.out.println("(g1 == g2) = " + (g1 == g2)); // true (Singleton)
        ctx.close();
    }

}
