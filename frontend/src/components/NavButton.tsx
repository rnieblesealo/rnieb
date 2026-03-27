import { useState } from "react"
import { Link } from "react-router-dom"

interface NavButtonProps {
  route: string,
  iconSrc: string
}

const NavButton = ({ route, iconSrc }: NavButtonProps) => {
  const [hovered, setHovered] = useState(false)

  return (
    <Link to={route}>
      <div
        onMouseEnter={() => setHovered(true)}
        onMouseLeave={() => setHovered(false)}
        className="relative overflow-hidden rounded-sm"
      >
        <img
          src={iconSrc}
          className="w-28 h-fit"
          style={{
            filter: hovered ? undefined : "saturate(0)"
          }}
        />
        {/* gradient overlay on top for ios 6 look */}
        <div className="bg-gradient-to-b from-white to-black opacity-30 w-full h-full absolute top-0" />
      </div>
    </Link>
  )
}

export default NavButton
