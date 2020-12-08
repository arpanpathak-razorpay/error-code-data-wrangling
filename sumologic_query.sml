"upi" and "GATEWAY_PAYMENT_VERIFY_RESPONSE"
| json auto nodrop
| count %payment.payment_id, 
%context.error.gateway_error_code, 
%context.error.gateway_error_desc, 
%context.error.internal_error_code, 
%context.response.data.gateway_response.status, 
%context.response.data.status, %timestamp
| sort by _count desc