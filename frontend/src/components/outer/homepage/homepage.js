import React from 'react';
import Navbar from "../navbar/navbar"
import HeroSection from './heroSection';
import CommentSection from './commentSection';

const Homepage = () => {
    return (
        <div>
            <Navbar />
            <HeroSection />
            <CommentSection />
        </div>
    )
}

export default Homepage;