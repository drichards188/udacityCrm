package models

type DetailResp struct {
	Status int
	Msg    string
}

type CustomerDb struct {
	count int
	db    map[int]Customer
}

func NewCustomerDb() *CustomerDb {
	db := &CustomerDb{count: 0, db: map[int]Customer{}}

	db.AddCustomer(Customer{Id: 0, Name: "david", Role: "dev", active: true})
	db.AddCustomer(Customer{Id: 0, Name: "allie", Role: "sales", active: true})
	db.AddCustomer(Customer{Id: 0, Name: "kovax", Role: "leadership", active: true})

	return db
}

func (cDB *CustomerDb) AddCustomer(newCustomer Customer) DetailResp {
	_, ok := cDB.db[newCustomer.Id]

	if ok {
		return DetailResp{Status: 0, Msg: "customer id already exists"}
	} else {
		cDB.count++
		newCustomer.Id = cDB.count
		cDB.db[cDB.count] = newCustomer
		return DetailResp{Status: 1, Msg: "ok"}
	}
}

func (cDB *CustomerDb) GetDb() map[int]Customer {
	return cDB.db
}

func (cDB *CustomerDb) GetCustomer(id int) Customer {
	return cDB.db[id]
}

func (cDB *CustomerDb) DeleteCustomer(id int) bool {
	delete(cDB.db, id)
	cDB.count--

	return true
}

func (cDB *CustomerDb) UpdateCustomer(customer Customer) Customer {
	cDB.db[customer.Id] = customer

	return customer
}
