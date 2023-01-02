import React, { useState } from 'react';
import {
  HeroSectionContainer,
  HeroImage,
  HeroTextContainer,
  HeroText,
  PenguinfyButtonContainer,
  PenguinfyButton,
  HeroDescriptionContainer
} from './HeroSectionElements';
import HeroImageMonday from '../../images/HeroImageMonday.png'
import { HeroMondayText } from '../../texts/HeroTexts';
import { HeroPenguinText } from '../../texts/HeroTexts';


const HeroSection = () => {
  const [isPenguinfied, setPenguinfied] = useState(true)

  return (
    <>
      <HeroSectionContainer>
        <HeroImage src={HeroImageMonday} />
        <HeroDescriptionContainer>
          <HeroTextContainer>
            <HeroText>{isPenguinfied ? HeroPenguinText : HeroMondayText}</HeroText>
          </HeroTextContainer>
          <PenguinfyButtonContainer>
            <PenguinfyButton onClick={() => setPenguinfied(!isPenguinfied)}>
              {isPenguinfied ? 'De-Penguinfy' : 'Penguinfy'}!
            </PenguinfyButton>
          </PenguinfyButtonContainer>
        </HeroDescriptionContainer>
      </HeroSectionContainer>
    </>
  )
}

export default HeroSection