import styled from 'styled-components'


export const HeroSectionContainer = styled.div`
    display: flex;
    height: 70vh;
`;

export const HeroImage = styled.img`
    width: 50vw;
`;

export const HeroDescriptionContainer = styled.div`
    background-color: #AEDBF0;
    width: 50vw;
    height: 70vh;
`;

export const HeroTextContainer = styled.div`
    display: flex;
    height: 50vh;
`;

export const HeroText = styled.div`
    display: flex;
    height: 40vh;
    margin: 5vh;
    text-align: center;
    align-items: center;
    font-size: 3rem;
`;

export const PenguinfyButtonContainer = styled.div`
    display: flex;
    height: 20vh;
    justify-content: flex-end;
`;

export const PenguinfyButton = styled.button`
    background-color: #CBF1FA;
    width: 20vw;
    height: 10vh;
    margin: 5vh;
    border: 0px;
    border-radius: 20px;
    font-size: 1.5rem;

    ${'' /* TODO: On hover, do something highlighting... */}
`;

