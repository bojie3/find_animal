import styled from 'styled-components';


export const Nav = styled.nav`
    height: 10vh;
    background-color: #3F7EB3;
`;

export const NavbarContainer = styled.div`
    align-items: center;
    justify-content: space-between;
    height: 10vh;
    display: flex;
`;

export const NavLogo = styled.img`
    height: 8vh;
    margin: 0px 0px 0px 3vh;
`;

export const NavButtonContainer = styled.div`
    display: flex;
`;

// TODO: Make this responsive
const NavButton = styled.button`
    font-size: 1rem;
    font-weight: bold;
    text-decoration: underline;
    color: #FFF;
    margin: 0px 2vw 0px 0px;
    border: 0px;
    border-radius: 20px;
    width: 9vw;
    height: 6vh;
    background-color: #6BA7CC;
`;

export const NavLoginButton = styled(NavButton)`
`;

export const NavRegisterButton = styled(NavButton)`
`;
