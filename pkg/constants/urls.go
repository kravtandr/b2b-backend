package constants

const LoginURL = "/user/login"
const RegisterURL = "/user/register"
const UserCheckEmailURL = "/user/check"
const ProfileURL = "/profile"
const LogoutURL = "/logout"
const UserInfoURL = "/user/{id}"

const FastOrderURL = "/fastOrder"

const FastRegisterURL = "/fastregister"
const CompanyURL = "/company/{id}"
const CompanyByInnFromDaDataURL = "/company/itn/{itn}"

const CategoryByIdURL = "/category"
const SearchCategoryURL = "/search/categories"

const ProductsListURL = "/search/products/?skip={skip}&limit={limit}"
const SearchProductsURL = "/search/products"

const CountryNameURL = "/country/{name}"
const CountryIdURL = "/country/id/{id}"
const CountryListURL = "/country/list"

const SightURL = "/sight/{id}"
const SightsByCountryURL = "/sights/{name}"
const SightSearch = "/sights/search"
const SightTag = "/sights/tag"
const Tags = "/tags"
const RandomTags = "/random/tags"

const ReviewAddURL = "/review"
const ReviewURL = "/review/{id}"

const TripPostURL = "/trip"
const TripURL = "/trip/{id}"
const SightsByTripURL = "/trip/sights/{id}"
const TripsByUserURL = "/trip/user"
const AddTripUserURL = "/trip/user/{id}"
const ShareTripURL = "/trip/share/{id}"
const SharedTripURL = "/trip/share/{code}/{id}"

const ShareTrip = "/trip/share/"

const UploadURL = "/upload"
const AvatarDirPath = "/avatars/"
const StaticServerURL = "http://194.58.104.204:3000"

const AlbumURL = "/album/{id}"
const AlbumAddURL = "/album"
const UploadAlbumPhotoURL = "/album/upload/{id}"
const AlbumsByUserURL = "/album/user"
