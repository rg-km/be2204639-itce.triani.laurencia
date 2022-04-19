package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository}
}

func (u *TransactionRepository) Pay(amount int) (int, error) {
<<<<<<< HEAD
<<<<<<< HEAD
	// TODO: replace this

	// selesaikan card dulu
	// total pricenya

	// kembalian = didapat dari amount - total 20.000 - 17.000 = 3.000

	total, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}

	return amount - total, nil
}
=======
	return 0, nil // TODO: replace this
}
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
	return 0, nil // TODO: replace this
}
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
