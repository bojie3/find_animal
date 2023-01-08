import React from 'react';
import {
  Nav,
  NavbarContainer,
  NavLogo,
  NavButtonContainer,
  NavLoginButton,
  NavRegisterButton
} from './NavbarElements';
import HomeIcon from '../../images/HomeIcon.png';



const Navbar = () => {
  return (
    <>
      <Nav>
        <NavbarContainer>
          <NavLogo src={HomeIcon} />
          <NavButtonContainer>
            <NavLoginButton>Login</NavLoginButton>
            <NavRegisterButton>Register</NavRegisterButton>
          </NavButtonContainer>
        </NavbarContainer>
      </Nav>
    </>

  )
}

export default Navbar