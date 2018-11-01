package model

import "time"

type Card struct {
	Id        int       `json:"Id"`
	CardNO    string    `json:"CardNO"`
	Status    int       `json:"Status"` // 是否使用(默认0未使用,1已使用,2已销毁)
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func (t *Card) TableName() string {
	return "gp_card"
}

type Charge struct {
	Id        int       `json:"Id"`
	Mobile    int       `json:"Mobile"`
	CardNO    string    `json:"CardNO"`
	Status    int       `json:"Status"` // 充值状态(默认0:待审核, 1已审核, 2审核不通过)
	Money     float64   `json:"Money"`  // 充值金额
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func (t *Charge) TableName() string {
	return "gp_charge"
}

type Guess struct {
	Id           int       `json:"Id"`
	Mobile       int       `json:"Mobile"`
	Project      string    `json:"Project"`      // 交易品种(上证指数, 恒生指数...)
	ProjectNO    string    `json:"ProjectNO"`    // 期数编号
	Deposit      float64   `json:"Deposit"`      // 下注金额
	DirectExpect string    `json:"DirectExpect"` // 交易方向(上涨:up, 下跌down)
	DirectReal   string    `json:"DirectReal"`   // 交易方向实际结果(上涨:up, 下跌:down)
	Result       int       `json:"Result"`       // 竞猜结果(默认0未知, 1赢, 2输)
	Bonus        float64   `json:"Bonus"`        // 竞猜奖金
	CreatedAt    time.Time `json:"CreatedAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
}

func (t *Guess) TableName() string {
	return "gp_guess"
}

type User struct {
	Id         int       `json:"Id"`
	Mobile     int       `json:"Mobile"`
	Email      string    `json:"Email"`
	Password   string    `json:"Password"`
	Money      float64   `json:"Money"`
	Frozen     float64   `json:"Frozen"`
	PayAccount string    `json:"PayAccount"`
	RealName   string    `json:"RealName"` // 真是姓名
	CreatedAt  time.Time `json:"CreatedAt"`
}

func (t *User) TableName() string {
	return "gp_user"
}

type Withdraw struct {
	Id        int       `json:"Id"`
	Mobile    int       `json:"Mobile"`
	Money     float64   `json:"Money"`    // 实际到账金额
	OrderNO   string    `json:"OrderNO"`  // 订单号
	Status    int       `json:"Status"`   // 提现状态(默认0:待审核, 1审核通过, 2审核不通过)
	MoneyAll  float64   `json:"MoneyAll"` // 提现总额
	Accrual   float64   `json:"Accrual"`  // 利息
	Mark      string    `json:"Mark"`     // 备注(未通过原因等)
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func (t *Withdraw) TableName() string {
	return "gp_withdraw"
}
