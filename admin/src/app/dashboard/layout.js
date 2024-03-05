import Secure from '@/components/Secure';
import SideTab from '@/components/SideTab';

export default function DashboardLayout({children}) {
    return (
        <div className='flex h-full min-h-screen'>
            {//<Secure />
            }
            <SideTab/>
            <main className='p-8 w-full overflow-auto'>
                {children}
            </main>
        </div>
    )
}