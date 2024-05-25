import {
  BadRequestException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { isValidObjectId, Model } from 'mongoose';
import * as bcrypt from 'bcrypt';
import { Admin } from './schemas/admin.schema';
import { CreateAdminDto } from './dto/create-admin.dto';
import { UpdateAdminDto } from './dto/update-admin.dto';
import { JwtService } from '@nestjs/jwt';

@Injectable()
export class AdminService {
  constructor(
    @InjectModel(Admin.name) private adminModel: Model<Admin>,
    private jwtService: JwtService,
  ) {}

  private validateSecret(secret: string): void {
    if (secret !== process.env.ADMIN_SECRET) {
      throw new BadRequestException('Invalid secret');
    }
  }

  async create(createAdminDto: CreateAdminDto): Promise<Admin> {
    this.validateSecret(createAdminDto.secret);

    if (await this.findByUsername(createAdminDto.username)) {
      throw new BadRequestException('Username already exists');
    }

    const hashedPassword = await bcrypt.hash(createAdminDto.password, 10);
    const createdAdmin = new this.adminModel({
      ...createAdminDto,
      password: hashedPassword,
    });
    const admin = await createdAdmin.save();

    admin.password = undefined;
    return admin;
  }

  async findAll(): Promise<Admin[]> {
    const admins = await this.adminModel.find().exec();
    return admins.map((admin) => {
      admin.password = undefined;
      return admin;
    });
  }

  async findOne(id: string): Promise<Admin> {
    this.validateObjectId(id);

    const admin = await this.adminModel.findById(id).exec();
    if (!admin) {
      throw new NotFoundException(`Admin with Id ${id} not found`);
    }

    admin.password = undefined;
    return admin;
  }

  async findByUsername(username: string): Promise<Admin> {
    const admin = await this.adminModel.findOne({ username }).exec();
    if (!admin) {
      throw new NotFoundException(`Admin with username ${username} not found`);
    }

    admin.password = undefined;
    return admin;
  }

  async getSessionData(token: string) {
    const decoded = this.jwtService.decode(token) as { sub: string };
    return this.adminModel.findById(decoded.sub).exec();
  }

  async remove(id: string, token: string): Promise<Admin> {
    const adminSession = await this.getSessionData(token);
    if (adminSession._id === id) {
      throw new BadRequestException('Cannot delete own account');
    }

    const admin = await this.adminModel.findByIdAndDelete(id).exec();
    if (!admin) {
      throw new NotFoundException(`Admin with Id ${id} not found`);
    }

    admin.password = undefined;
    return admin;
  }

  async update(id: string, updateAdminDto: UpdateAdminDto): Promise<Admin> {
    if (updateAdminDto.password) {
      updateAdminDto.password = await bcrypt.hash(updateAdminDto.password, 10);
    } else {
      delete updateAdminDto.password;
    }

    this.validateObjectId(id);

    const admin = await this.adminModel
      .findByIdAndUpdate(id, updateAdminDto, { new: true })
      .exec();
    if (!admin) {
      throw new NotFoundException(`Admin with Id ${id} not found`);
    }

    admin.password = undefined;
    return admin;
  }

  private validateObjectId(id: string): void {
    if (!isValidObjectId(id)) {
      throw new BadRequestException(`Invalid Id format: ${id}`);
    }
  }
}
