package database

var (
	promoTable = `
		CREATE TABLE IF NOT EXISTS public.promo (
			promo_id uuid NOT NULL,
			"promo_attr" varchar NOT NULL,
			"promo_code" varchar NOT NULL UNIQUE,
			CONSTRAINT promo_pk PRIMARY KEY (promo_id)
		);
		`

	orderTable = `
		CREATE TABLE IF NOT EXISTS public.orders (
			order_id uuid NOT NULL,
			"order_attr" varchar NULL,
			"promo_code" varchar NULL,
			CONSTRAINT orders_pk PRIMARY KEY (order_id),
			FOREIGN KEY (promo_code) REFERENCES promo (promo_code)
		);
		`

	orderPromoTable = `
		CREATE TABLE IF NOT EXISTS public.order_promo (
			id uuid NOT NULL,
			order_id uuid NOT NULL,
			promo_id uuid NOT NULL,
			CONSTRAINT order_promo_pk PRIMARY KEY (id),
			FOREIGN KEY (order_id) REFERENCES orders (order_id),
			FOREIGN KEY (promo_id) REFERENCES promo (promo_id)
		);
		`
)
