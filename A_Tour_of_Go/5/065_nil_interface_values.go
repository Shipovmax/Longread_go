package atourofgo

func nil_interface_values() {

	var i I

	describe(i) // должно вывести: (<nil>, <nil>)

	i.M()
}
