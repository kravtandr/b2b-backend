package constants

// Auth
const LoginURL = "/user/login"
const RegisterURL = "/user/register"
const UserCheckEmailURL = "/user/check"
const ProfileURL = "/profile"
const LogoutURL = "/logout"
const UserInfoURL = "/user/{id}"
const UserInfoByCookieURL = "/user/me"

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
const SearchProductsURL = "/search/products"

// Chat
const CheckIfUniqChat = "/chat/unique"
const InitChat = "/chat/init"
const ProductChatURL = "/ws"
const AllChats = "/allchats"
const AllMsgsFromChat = "/chat/{id}"

// Healthchecks
const AuthHealthCheck = "/auth/health/live"
const ChatHealthCheck = "/chat/health/live"
const CompanyHealthCheck = "/company/health/live"
const FastOrderHealthCheck = "/fo/health/live"
const ProductsCategoriesHealthCheck = "/pc/health/live"
