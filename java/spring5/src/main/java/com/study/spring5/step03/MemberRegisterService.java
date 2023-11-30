package com.study.spring5.step03;

/**
 * 의존 관계
 * 한 클래스가 다른 클래스의 매서드를 실행할 때 이를 '의존한다.' 고한다.
 * ->'MemberRegisterService는 MemberDao객체의 메서드를 실행한다.'
 * -> MemberRegisterService는 MemberDao에 의존한다.
 *
 * */
public class MemberRegisterService {

    private MemberDao memberDao = new MemberDao();


}
