package com.study.spring5.step02;

import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.GenericGroovyApplicationContext;
import org.springframework.context.support.GenericXmlApplicationContext;

public class Step02Application {

    public static void main(String[] args) {

        /**
         * ApplicationContext 상속 관계
         * | BeanFactory
         * |- ListableBeanFactory
         * |-- ApplicationContext
         * |--- ConfigurableApplicationContext
         * |---- AbstractApplicationContext
         * |---- BeanDefinitionRegistry
         * |----- GenericApplicationContext
         * |----- AnnotationConfigRegistry
         * |------ AnnotationConfigApplicationContext (.java configuration)
         * |------ GenericXmlApplicationContext (XML file configuration)
         * |------ GenericGroovyApplicationContext (Groovy configuration)
         * */
        AnnotationConfigApplicationContext ctx =
                new AnnotationConfigApplicationContext(AppContext.class);

        Greeter g = ctx.getBean("greeter", Greeter.class);
        String msg = g.greet("스프링");
        System.out.println(msg);
        ctx.close();
    }
}
