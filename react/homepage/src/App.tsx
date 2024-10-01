import Icon from "@mdi/react";
import {
  mdiGithub,
  mdiLinkedin,
  mdiTwitter,
  mdiArrowTopRightThinCircleOutline,
  mdiEmailOutline,
  mdiPhoneDialOutline,
} from "@mdi/js";
import "./index.css";

function App() {
  return (
    <div>
      <div className="header-background">
        <div className="header-container">
          <img className="header-picture" src="/smiling.jpg" alt="" />
          <div className="header-card">
            <div className="header-card-title">About Me</div>
            <div className="header-card-body">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptas
              quia voluptatibus nam necessitatibus unde quos veritatis aperiam
              quo nesciunt magni. Dolor, dignissimos dicta architecto cumque
              dolorum officia quos aut aliquam! Lorem ipsum dolor sit amet
              consectetur adipisicing elit. Animi incidunt veniam ad minima
              nihil commodi vel atque placeat ipsum? Ut ipsam eius accusantium
              sint ea beatae illum molestiae illo in.
            </div>
            <div className="header-card-icons">
              <Icon className="social-icons" path={mdiLinkedin} size={2} />
              <Icon className="social-icons" path={mdiGithub} size={2} />
              <Icon className="social-icons" path={mdiTwitter} size={2} />
            </div>
          </div>
        </div>
      </div>
      <div className="body-background">
        <div className="body-title">My Work</div>
        <div className="body-cards-container">
          <div className="body-cards-card">
            <div className="body-cards-card-top">
              <div className="body-cards-card-picture">
                Screenshot of Project
              </div>
            </div>
            <div className="body-cards-card-body">
              <div className="body-cards-card-body-header">
                <div className="body-cards-card-body-title">Project Name</div>
                <div className="body-cards-card-body-icons">
                  <Icon path={mdiGithub} size={1} />
                  <Icon path={mdiArrowTopRightThinCircleOutline} size={1} />
                </div>
              </div>
              <div className="body-cards-card-body-description">
                Short description of the project. Just a couple of sentences
                will do
              </div>
            </div>
          </div>
          <div className="body-cards-card">
            <div className="body-cards-card-top">
              <div className="body-cards-card-picture">
                Screenshot of Project
              </div>
            </div>
            <div className="body-cards-card-body">
              <div className="body-cards-card-body-header">
                <div className="body-cards-card-body-title">Project Name</div>
                <div className="body-cards-card-body-icons">
                  <Icon path={mdiGithub} size={1} />
                  <Icon path={mdiArrowTopRightThinCircleOutline} size={1} />
                </div>
              </div>
              <div className="body-cards-card-body-description">
                Short description of the project. Just a couple of sentences
                will do
              </div>
            </div>
          </div>
          <div className="body-cards-card">
            <div className="body-cards-card-top">
              <div className="body-cards-card-picture">
                Screenshot of Project
              </div>
            </div>
            <div className="body-cards-card-body">
              <div className="body-cards-card-body-header">
                <div className="body-cards-card-body-title">Project Name</div>
                <div className="body-cards-card-body-icons">
                  <Icon path={mdiGithub} size={1} />
                  <Icon path={mdiArrowTopRightThinCircleOutline} size={1} />
                </div>
              </div>
              <div className="body-cards-card-body-description">
                Short description of the project. Just a couple of sentences
                will do
              </div>
            </div>
          </div>
          <div className="body-cards-card">
            <div className="body-cards-card-top">
              <div className="body-cards-card-picture">
                Screenshot of Project
              </div>
            </div>
            <div className="body-cards-card-body">
              <div className="body-cards-card-body-header">
                <div className="body-cards-card-body-title">Project Name</div>
                <div className="body-cards-card-body-icons">
                  <Icon path={mdiGithub} size={1} />
                  <Icon path={mdiArrowTopRightThinCircleOutline} size={1} />
                </div>
              </div>
              <div className="body-cards-card-body-description">
                Short description of the project. Just a couple of sentences
                will do
              </div>
            </div>
          </div>
          <div className="body-cards-card">
            <div className="body-cards-card-top">
              <div className="body-cards-card-picture">
                Screenshot of Project
              </div>
            </div>
            <div className="body-cards-card-body">
              <div className="body-cards-card-body-header">
                <div className="body-cards-card-body-title">Project Name</div>
                <div className="body-cards-card-body-icons">
                  <Icon path={mdiGithub} size={1} />
                  <Icon path={mdiArrowTopRightThinCircleOutline} size={1} />
                </div>
              </div>
              <div className="body-cards-card-body-description">
                Short description of the project. Just a couple of sentences
                will do
              </div>
            </div>
          </div>
          <div className="body-cards-card">
            <div className="body-cards-card-top">
              <div className="body-cards-card-picture">
                Screenshot of Project
              </div>
            </div>
            <div className="body-cards-card-body">
              <div className="body-cards-card-body-header">
                <div className="body-cards-card-body-title">Project Name</div>
                <div className="body-cards-card-body-icons">
                  <Icon path={mdiGithub} size={1} />
                  <Icon path={mdiArrowTopRightThinCircleOutline} size={1} />
                </div>
              </div>
              <div className="body-cards-card-body-description">
                Short description of the project. Just a couple of sentences
                will do
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="footer-background">
        <div className="footer-container">
          <div className="footer-contact-container">
            <div className="footer-contact-title">Contact Me</div>
            <div className="footer-contact-message">
              Please get in touch if you think our work could be mutually
              beneficial!
            </div>
            <div className="footer-contact-message">
              1234 Random Road Random Town, California 12345
            </div>
            <div className="footer-contact-data">
              <div>
                <div className="footer-contact-data-container">
                  <Icon path={mdiPhoneDialOutline} size={2} />
                  <span>555-555-5555</span>
                </div>
                <div className="footer-contact-data-container">
                  <Icon path={mdiEmailOutline} size={2} />
                  <span>ashleywilliams.is.not.real@gmail.com</span>
                </div>
              </div>
            </div>
            <div className="footer-contact-icons">
              <Icon className="footer-social-icons" path={mdiGithub} size={2} />
              <Icon
                className="footer-social-icons"
                path={mdiLinkedin}
                size={2}
              />
              <Icon
                className="footer-social-icons"
                path={mdiTwitter}
                size={2}
              />
            </div>
          </div>

          <img className="footer-picture" src="/smiling.jpg" alt="" />
        </div>
      </div>
    </div>
  );
}

export default App;
