context: Order-Taking
// ----------------------
// Simple types
// ----------------------

// Product codes

data ProductCode = WidgetCode OR GizmoCode
data WidgetCode = string starting with "W" then 4 digits
data GizmoCode = ...

// Order Quantity
data OrderQuantity = UnitQuantity OR KilogramQuantity
data UnitQuantity = ... 
data KilogramQuantity = ...

// ----------------------
// Order life cycle
// ----------------------

// ----- unvalidated state -----
data UnvalidatedOrder =
    UnvalidatedCustomerInfo
    AND UnvalidatedShippingAddress
    AND UnvalidatedBillingAddress
    AND list of UnvalidatedOrderLine

data UnvalidatedOrderLine =
    UnvalidatedProductCode
    AND UnvalidatedOrderQuantity
    
// ----- validated state -----
data ValidatedOrder = ...
data ValidatedOrderLine = ...

// ----- priced state -----
data PricedOrder = ...
data PricedOrderLine = ...

// ----- output events -----
data OrderAcknowledgmentSent = ...
data OrderPlaced = ...
data BillableOrderPlaced = ...

// ----------------------
// Workflows
// ----------------------
workflow "Place Order" =
    input: UnvalidatedOrder
    output (on success):
        OrderAcknowledgmentSent
        AND OrderPlaced (to send to shipping)
        AND BillableOrderPlaced (to send to billing)
    output (on error):
        InvalidOrder
        
//etc