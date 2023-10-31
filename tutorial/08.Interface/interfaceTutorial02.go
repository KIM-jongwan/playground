package main

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}

	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber, s.cost()
	}

	return "", 0.0

}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.1
	}
	return float64(len(e.body)) * 0.01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	}
	return float64(len(s.body)) * 0.03
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func main() {

}
