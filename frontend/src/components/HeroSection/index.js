import React, { useState } from 'react'
import Video from '../../videos/video.mp4'
import { Button } from '../ButtonElement'
import {
  HeroContainer,
  HeroBg,
  VideoBg,
  HeroContent,
  HeroH1,
  HeroP,
  HeroBtnWrapper
} from './HeroElements'

const HeroSection = () => {
  const [toggle, setToggle] = useState(false)

  const onToggle = () => {
    setToggle(!toggle)
  }

  return (
    <HeroContainer>
      <HeroBg>
        <VideoBg autoPlay loop muted src={Video} type='video/mp4' />
      </HeroBg>
      <HeroContent>
        <HeroH1>My Little Penguins</HeroH1>
        <HeroP>
          HONK HONK HONK
        </HeroP>
        <HeroBtnWrapper>
          <Button onClick={onToggle}>
            {toggle ? "DePenguinfy" : "Penguinfy"}!
          </Button>
        </HeroBtnWrapper>
      </HeroContent>
    </HeroContainer>
  )
}

export default HeroSection