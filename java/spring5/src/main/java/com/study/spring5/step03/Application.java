package com.study.spring5.step03;

public class Application {

    public static void main(String[] args) {
        /*MemberDao 객체를 생성하여 주입*/
        MemberDao memberDao = new MemberDao();
        MemberRegisterService memberRegisterService = new MemberRegisterService(memberDao);
        ChangePasswordService changePasswordService = new ChangePasswordService();
        changePasswordService.setMemberDao(memberDao);
    }
}
