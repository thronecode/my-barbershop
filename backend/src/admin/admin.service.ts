import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import * as bcrypt from 'bcrypt';
import { Admin } from './schemas/admin.schema';
import { CreateAdminDto } from './dto/create-admin.dto';
import { UpdateAdminDto } from './dto/update-admin.dto';

@Injectable()
export class AdminService {
  constructor(@InjectModel(Admin.name) private adminModel: Model<Admin>) {}

  async create(createAdminDto: CreateAdminDto): Promise<Admin> {
    const hashedPassword = await bcrypt.hash(createAdminDto.password, 10);
    const createdAdmin = new this.adminModel({
      ...createAdminDto,
      password: hashedPassword,
    });
    return createdAdmin.save();
  }

  async findAll(): Promise<Admin[]> {
    return this.adminModel.find().exec();
  }

  async findOne(id: string): Promise<Admin> {
    return this.adminModel.findById(id).exec();
  }

  async findByUsername(username: string): Promise<Admin> {
    return this.adminModel.findOne({
      username,
    });
  }

  async remove(id: string): Promise<Admin> {
    return this.adminModel.findByIdAndDelete(id).exec();
  }

  async update(id: string, updateAdminDto: UpdateAdminDto): Promise<Admin> {
    const hashedPassword = await bcrypt.hash(updateAdminDto.password, 10);
    return this.adminModel
      .findByIdAndUpdate(
        id,
        { ...updateAdminDto, password: hashedPassword },
        { new: true },
      )
      .exec();
  }
}
