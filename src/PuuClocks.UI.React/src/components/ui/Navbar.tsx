import { useTranslation } from 'react-i18next';
import { NavLink } from 'react-router-dom';
import logo from '~/assets/images/PuuClocksLogoAlpha.png';
import { Button } from './Button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from '~/components/ui/DropdownMenu';
import { ModeToggle } from './ModeToggle';

export const Navbar = () => {
  const { t, i18n } = useTranslation('common');

  const changeLanguage = (lng: string) => {
    i18n.changeLanguage(lng);
  };

  const getLabel = (language: string) => {
    switch (language) {
      case 'en-US':
        return 'English';
      case 'pl-PL':
        return 'Polski';
    }
  };

  return (
    <div className="w-full flex items-center justify-between">
      <NavLink id="home" to="/">
        <img className="max-w-16" src={logo} alt="PuuClocksLogo" />
      </NavLink>
      <div className="flex items-center gap-2">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline">{getLabel(i18n.language)}</Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-32">
            <DropdownMenuLabel>
              {t('AppSettings.SelectLanguage')}
            </DropdownMenuLabel>
            <DropdownMenuItem>
              <Button variant="outline" onClick={() => changeLanguage('en-US')}>
                English
              </Button>
            </DropdownMenuItem>
            <DropdownMenuItem>
              <Button variant="outline" onClick={() => changeLanguage('pl-PL')}>
                Polski
              </Button>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
        <ModeToggle />
      </div>
    </div>
  );
};
