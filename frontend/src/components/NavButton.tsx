import { useState } from "react"
import { NavLink } from "react-router-dom"

interface NavButtonProps {
  route: string,
  iconSrc: string
}

const NavButton = ({ route, iconSrc }: NavButtonProps) => {
  const [hovered, setHovered] = useState(false)

  return (
    <NavLink to={route}
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
      className="relative overflow-hidden rounded-sm"
    >
      {/* isactive gets passed down by navlink to everything inside the ( );
          this is true whenever this is the active route */}

      {/* navlink accepts a function with isactive param as its children prop
          isactive is passed by navlink to the stuff inside the () */ }

      {({ isActive }) => (
        <>
          <img
            src={iconSrc}
            className="w-28"
            style={{
              filter: (hovered || isActive) ? undefined : "saturate(0)"
            }}
          />
          {/* gradient overlay on top for ios 6 look */}
          {isActive ?
            <div className="bg-gradient-to-t from-white to-black opacity-30 w-full h-full absolute top-0" /> :
            <div className="bg-gradient-to-b from-white to-black opacity-30 w-full h-full absolute top-0" />
          }
        </>
      )}
    </NavLink>
  )
}

export default NavButton
