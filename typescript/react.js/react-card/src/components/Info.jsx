import React from 'react'
import { Button } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faLinkedin } from '@fortawesome/free-brands-svg-icons'
import { faEnvelope } from '@fortawesome/free-solid-svg-icons'

function About() {
  return (
    <div className='info'>
      <h2 className='info--name'>Corina Pulido</h2>
      <h3 className='info--jobtitle'>Lead Planner</h3>
      <h5 className='info--website'>Corina Website</h5>
      <Button className='info-emailButton' variant="contained">
        <FontAwesomeIcon
          className="info--icons"
          icon={faEnvelope}
        />{' '}
        Email
      </Button>
      <Button className='info--linkedinButton' variant="contained">
        <FontAwesomeIcon
          className="info--icons"
          icon={faLinkedin}
        />{' '}
        LinkedIn
      </Button>
    </div>
  )
}

export default About
