
CREATE TABLE public.staff (
	id serial4 NOT NULL,
	name varchar NULL,
	email varchar NULL,
	phone varchar NULL,
	entrytime int8 NOT NULL
);
CREATE UNIQUE INDEX staff_email_idx ON public.staff USING btree (email, phone);



INSERT INTO public.staff (name,email,phone,entrytime) VALUES
	 ('Janet Billyskid','sirpros3@gmail.com','080621401039',1642562202),
	 ('Johnet Bosco','sirprosxxx@gmail.com','08012521401039',1642568313);