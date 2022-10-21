import std/httpclient

#[
check
  warming-up
    points 5
    timeout 2s
    - send
      expect
      errormsg "think twice and try again"
  common-params "占30分，这里是章节测试的注释"
    points 30, 
    - send 
      points 10 # 这里的分数不能超过上面的分数，总分为100，这些都需要宏来检测
      expect
      errormsg ""
    - send 
      expect
      errormsg
    - send 
      expect
      errormsg ""
      

  middlewire 20
    points 30, 

  authority 30
    points 30, 
    - send
        
        # 这里参考 apifox
      expect
        # 这里也参考 apifox
      errormsg ""

  custom-server
    points 30, 

  configuration
    points 30, 

  project-layer
    points 30, 

  template-mvc
    points 30, 

  dependency-inject
    points 30, 

]#
