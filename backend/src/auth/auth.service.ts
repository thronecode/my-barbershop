import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import * as bcrypt from 'bcrypt';
import { Admin } from '../admin/schemas/admin.schema';
import { LoginDto } from './dto/login.dto';

@Injectable()
export class AuthService {
  constructor(
    @InjectModel(Admin.name) private adminModel: Model<Admin>,
    private jwtService: JwtService,
  ) {}

  async validateAdmin(loginDto: LoginDto): Promise<Admin | null> {
    const admin = await this.adminModel
      .findOne({ username: loginDto.username })
      .exec();
    if (admin && (await bcrypt.compare(loginDto.password, admin.password))) {
      return admin;
    }
    return null;
  }

  // async getSessionData(token: string) {
  //   const decoded = this.jwtService.decode(token) as { sub: string };
  //   return this.adminModel.findById(decoded.sub).exec();
  // }

  async login(admin: Admin) {
    const payload = { username: admin.username, sub: admin._id };
    return {
      access_token: this.jwtService.sign(payload, {
        secret: process.env.JWT_SECRET,
        expiresIn: '7d',
      }),
    };
  }
}
