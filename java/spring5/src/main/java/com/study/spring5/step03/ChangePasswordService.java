package com.study.spring5.step03;

public class ChangePasswordService {
    private MemberDao memberDao;
    public ChangePasswordService(MemberDao memberDao){
        this.memberDao = memberDao;
    }
}
