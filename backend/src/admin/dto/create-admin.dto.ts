import { IsString, IsNotEmpty, IsStrongPassword } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class CreateAdminDto {
  @ApiProperty({ description: 'Username of the admin' })
  @IsString()
  @IsNotEmpty()
  readonly username: string;

  @ApiProperty({ description: 'Password of the admin' })
  @IsStrongPassword()
  @IsNotEmpty()
  readonly password: string;

  @ApiProperty({ description: 'Secret key of the admin' })
  @IsString()
  @IsNotEmpty()
  readonly secret: string;
}
