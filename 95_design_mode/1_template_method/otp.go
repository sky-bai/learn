package main

// 模板方法

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

// Otp 该结构体实现某一个接口
type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

// 让我们来考虑一个一次性密码功能 （OTP） 的例子。 将 OTP 传递给用户的方式多种多样 （短信、 邮件等）。 但无论是短信还是邮件， 整个 OTP 流程都是相同的：
//
//生成随机的 n 位数字。
//在缓存中保存这组数字以便进行后续验证。
//准备内容。
//发送通知。
//后续引入的任何新 OTP 类型都很有可能需要进行相同的上述步骤。
//
//因此， 我们会有这样的一个场景， 其中某个特定操作的步骤是相同的， 但实现方式却可能有所不同。 这正是适合考虑使用模板方法模式的情况。
//
//首先， 我们定义一个由固定数量的方法组成的基础模板算法。 这就是我们的模板方法。 然后我们将实现每一个步骤方法， 但不会改变模板方法。

// 使用方法：
// 1.使用接口抽象需要走的流程
// 2.使用上层结构体对象包裹该接口并实现对应接口 定义模版方法类
// 3.构造下层具体实现类进行具体接口实现 该具体实现类要组合上层结构体
