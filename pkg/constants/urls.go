package constants

// Auth
const LoginURL = "/user/login"
const RegisterURL = "/user/register"
const UserCheckEmailURL = "/user/check"
const ProfileURL = "/profile"
const LogoutURL = "/logout"
const UserInfoURL = "/user/{id}"
const UserInfoByCookieURL = "/user/me"

// yookassa
const CreatePaymentAddBalanceURL = "/payment/addbalance"
const CheckPaymentURL = "/payment/check"
const HandlePaidPaymentsURL = "/user/payment/handle/all"
const GetUsersPaymentsURL = "/user/payments/list"
const CountUsersPaymentsURL = "/user/payments/count"

// FastOrder
const FastOrderURL = "/fastOrder"
const LandingOrderURL = "/landingOrder"

// Company
const FastRegisterURL = "/fastregister"
const CompanyURL = "/company/{id}"
const CompanyByInnFromDaDataURL = "/company/itn/{itn}"

// Categories
const CategoryByIdURL = "/category"
const SearchCategoryURL = "/search/categories"

// Products
const ProductURL = "/product/{id}"
const AddProductURL = "/product/add"
const ProductsListURL = "/products/list"
const ProductsListByFiltersURL = "/products/list/filter" // {category_name}
const SearchProductsURL = "/search/products"
const CompanyProductsListURL = "/company/products"
const UpdateProductURL = "/product/edit"

// Chat
const CheckIfUniqChat = "/chat/unique"
const InitChat = "/chat/init"
const DeleteChat = "/chat/delete"
const ChatChangeStatus = "/chat/status"
const ProductChatURL = "/ws"
const AllChats = "/allchats"
const AllMsgsFromChat = "/chat/{id}"

// Healthchecks
const AuthHealthCheck = "/auth/health/live"
const ChatHealthCheck = "/chat/health/live"
const CompanyHealthCheck = "/company/health/live"
const FastOrderHealthCheck = "/fo/health/live"
const ProductsCategoriesHealthCheck = "/pc/health/live"
