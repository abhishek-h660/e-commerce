import './../style/homeHeader.css'
import SearchBox from './SearchBar';

function HomeHeader() {
    return(
        <div className="home-header">
            <div className='header-item' id='company-logo'><img src='/companyLogo.png' alt='Company logo' /></div>
            <SearchBox />
            <div className='header-item'>
                <div className='header-item-child' id="seller">Become a Seller</div>
                <div className='header-item-child' id="more">More</div>
                <div className='header-item-child' id="cart">Cart</div>
            </div>
        </div>
    )
}

export default HomeHeader;