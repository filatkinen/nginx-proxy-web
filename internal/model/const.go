package model

const (
	PaymentTypeCard PaymentType = 1
	PaymentTypeSBP  PaymentType = 2

	NotificationTypeEmail NotificationType = 1
	NotificationTypeSMS   NotificationType = 1

	OrderStatusNew     OrderStatus = 1
	OrderStatusProcess OrderStatus = 2
	OrderStatusOk      OrderStatus = 3
	OrderStatusNotOk   OrderStatus = 4
	OrderStatusCancel  OrderStatus = 4

	ServiceStatusAk    ServiceStatus = 1
	ServiceStatusOK    ServiceStatus = 2
	ServiceStatusNotOK ServiceStatus = 2

	ServiceIDWarehouse    ServiceID = 1
	ServiceIDPayment      ServiceID = 2
	ServiceIDNotification ServiceID = 3

	InvoiceTypeIn      InvoiceType = 1
	InvoiceTypeOut     InvoiceType = 2
	InvoiceTypeRetract InvoiceType = 3
)
