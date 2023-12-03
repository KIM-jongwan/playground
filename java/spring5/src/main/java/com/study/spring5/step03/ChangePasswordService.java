package com.study.spring5.step03;

import com.study.spring5.common.Member;
import com.study.spring5.common.MemberNotFoundException;

public class ChangePasswordService {
    private MemberDao memberDao;
    /*public ChangePasswordService(MemberDao memberDao){
        this.memberDao = memberDao;
    }*/

    public void changePassword(String email, String oldPwd, String newPwd){
        Member member = memberDao.selectByEmail(email);
        if(member == null){
            throw new MemberNotFoundException();
        }
        member.changePassword(oldPwd, newPwd);
        memberDao.update(member);
    }

    public void setMemberDao(MemberDao memberDao){
        this.memberDao = memberDao;
    }
}
