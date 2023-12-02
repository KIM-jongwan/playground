package com.study.spring5.step03;

import com.study.spring5.common.Member;
import com.study.spring5.common.RegisterRequest;
import com.sun.jdi.request.DuplicateRequestException;


/**
 * 의존 관계
 * 한 클래스가 다른 클래스의 매서드를 실행할 때 이를 '의존한다.' 고한다.
 * ->'MemberRegisterService는 MemberDao객체의 메서드를 실행한다.'
 * -> MemberRegisterService는 MemberDao에 의존한다.
 *
 * */
public class MemberRegisterService {

    /*
    1. 의존 객체 직접 생성
    private MemberDao memberDao = new MemberDao(); */

    /*
    * 2. 의존객체를 전달받아 생성
    * */
    private MemberDao memberDao;
    public  MemberRegisterService(MemberDao memberDao){
        this.memberDao = memberDao;
    }


    public void regist(RegisterRequest req){
        Member member = memberDao.selectByEmail(req.getEmail());

        if(member != null){
            //exception email duplication
            throw new DuplicateRequestException("email duplication");
        }

        Member newMember = new Member(req.getEmail(), req.getName(),req.getPassword());
        memberDao.insert(newMember);

    }


}
