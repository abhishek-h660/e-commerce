import HomeHeader from './header'
import { useEffect } from 'react';
const CartItems = () => {
    useEffect(()=>{
        const skip = 0;
        const limit = 20;
        fetch(`http://localhost:8080/list_cart_items?skip=${skip}&limit=${limit}`, {
            method: "GET"
        }).then(res => {
            return res.json
        }).then(data => {
            console.log(data)
        })
    })

    return(
        <div className="cart-container">
            <HomeHeader />
            These are Cart Items
        </div>
    )
}

export default CartItems;