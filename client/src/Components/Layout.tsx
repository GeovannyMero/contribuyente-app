import { Outlet } from "react-router";

const Layout = () => {
    return (
        <div>
            <header>
                <div className="navbar bg-base-100 shadow-sm">
                    <div className="flex-1">
                        <a className="btn btn-ghost text-xl">Contribuyentes</a>
                    </div>
                    <div className="flex gap-2">
                        <input type="text" placeholder="Buscar" className="input input-bordered w-24 md:w-auto" />
                        <div className="dropdown dropdown-end">
                            <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
                                <div className="w-10 rounded-full">
                                    <img
                                        alt="User"
                                        src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </header>

            {/* 2. El Contenido de la Ruta Hija se renderizará aquí */}
            <main style={{ padding: '20px' }}>
                <Outlet />
            </main>

            {/* 3. Elementos Estáticos (Se repiten en todas las rutas) */}
            <footer style={{ background: '#ccc', padding: '10px', textAlign: 'center' }}>
                &copy; {new Date().getFullYear()} Contribuyentes
            </footer>
        </div>

    )
}

export default Layout;